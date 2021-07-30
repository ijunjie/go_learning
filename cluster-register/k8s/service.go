package k8s

import (
	"cluster-register/infra"
	"context"
	"encoding/base64"
	"flag"
	"fmt"
	v1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"math"
	"path/filepath"
	"time"
)

const (
	clusterRoleBindingName = "admin-service-account"
)

type K8sInfoRequest struct {
	ClusterName string
	K8sType     string
}

type K8sInfoResult struct {
	ClusterName string
	Host        string
	BasicKey    string
	Cpu         int
	Cu          int
	MemGB       int
	Type        string
	Env         int
}

func (info *K8sInfoResult) ToClusterConfigInsert() *infra.ClusterConfigInsert {
	return &infra.ClusterConfigInsert{
		ClusterName:    info.ClusterName,
		Host:           info.Host,
		RootCuNum:      info.Cu,
		BasicKey:       info.BasicKey,
		RmHost:         "",
		NmHost:         "",
		ClusterType:    info.Env,
		ClusterKind:    1,
		HadoopMasterIp: "",
	}
}

func K8sInfo(request *K8sInfoRequest, timeoutSeconds int) (*K8sInfoResult, error) {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err0 := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err0 != nil {
		return nil, err0
	}

	config.Timeout = time.Duration(timeoutSeconds) * time.Second

	// create the clientset
	clientset, err1 := kubernetes.NewForConfig(config)
	if err1 != nil {
		return nil, err1
	}

	_, err2 := clientset.RbacV1().ClusterRoleBindings().Get(context.TODO(), clusterRoleBindingName, metav1.GetOptions{})
	if errors.IsNotFound(err2) {
		log.Printf("create clusterrolebinding %s ...\n", clusterRoleBindingName)
		// 不存在则创建
		_, err3 := clientset.RbacV1().ClusterRoleBindings().Create(context.TODO(),
			&v1.ClusterRoleBinding{
				TypeMeta: metav1.TypeMeta{
					Kind:       "ClusterRoleBinding",
					APIVersion: "rbac.authorization.k8s.io/v1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name: clusterRoleBindingName,
				},
				Subjects: []v1.Subject{{
					Kind:      "ServiceAccount",
					Name:      "default",
					Namespace: "kube-system",
				}},
				RoleRef: v1.RoleRef{
					APIGroup: "rbac.authorization.k8s.io",
					Kind:     "ClusterRole",
					Name:     "cluster-admin",
				},
			},
			metav1.CreateOptions{
				TypeMeta:     metav1.TypeMeta{},
				DryRun:       nil,
				FieldManager: "",
			})
		if err3 != nil {
			return nil, err3
		}
	} else {
		log.Printf("%s already exists \n", clusterRoleBindingName)
	}

	log.Printf("MASTER_URL=%s \n", config.Host)

	// kubectl describe sa default -n kube-system
	sa, errSa := clientset.CoreV1().ServiceAccounts("kube-system").Get(context.TODO(), "default",
		metav1.GetOptions{})
	if errSa != nil {
		return nil, errSa
	}
	tokenName := sa.Secrets[0].Name
	log.Printf("TOKEN_NAME=%s", tokenName)

	// kubectl get secret default-token-hw9dn -n kube-system -o yaml
	secret, errSecret := clientset.CoreV1().Secrets("kube-system").Get(context.TODO(), tokenName, metav1.GetOptions{})
	if errSecret != nil {
		return nil, errSecret
	}
	caCrt := base64.StdEncoding.EncodeToString(secret.Data["ca.crt"])
	token := base64.StdEncoding.EncodeToString(secret.Data["token"])
	//log.Printf("CACRT=%s\n", caCrt)
	//log.Printf("TOKEN=%s\n", token)

	list, err4 := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{
		FieldSelector: "spec.unschedulable=false",
	})
	if err4 != nil {
		return nil, err4
	}

	quantityCpu := resource.NewQuantity(0, resource.BinarySI)
	quantityMem := resource.NewQuantity(0, resource.BinarySI)
	for _, value := range list.Items {
		quantityCpu.Add(*value.Status.Allocatable.Cpu())
		quantityMem.Add(*value.Status.Allocatable.Memory())
	}

	totalCpu := quantityCpu.Value()
	fmt.Printf("total cpu value= %v (%v)\n", totalCpu, quantityCpu.Format)
	fmt.Printf("total mem bytes= %v (%v)\n", quantityMem.Value(), quantityMem.Format)
	totalMemGi := int64(float64(quantityMem.Value()) / float64(1024*1024*1024))
	fmt.Printf("total mem Gi= %v\n", totalMemGi)

	totalCu := int64(math.Min(float64(totalCpu), float64(totalMemGi/4)))
	fmt.Printf("total Cu=%v\n", totalCu)

	// '{\"ca.crt\":"$CA_CRT",\"token\":"$DATA_TOKEN"}'
	basicKey := fmt.Sprintf("{\"ca.crt\":\"%s\",\"token\":\"%s\"}", caCrt, token)

	return &K8sInfoResult{
		ClusterName: request.ClusterName,
		Host:        config.Host,
		BasicKey:    basicKey,
		Cpu:         int(totalCpu),
		Cu:          int(totalCu),
		MemGB:       int(totalMemGi),
		Type:        request.K8sType,
		Env:         infra.TypeNumber(request.K8sType),
	}, nil
}
