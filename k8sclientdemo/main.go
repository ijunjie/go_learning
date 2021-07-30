package main

import (
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
)

const (
	clusterRoleBindingName = "admin-service-account"
)

func main() {
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
		panic(err0.Error())
	}

	// create the clientset
	clientset, err1 := kubernetes.NewForConfig(config)
	if err1 != nil {
		panic(err1.Error())
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
			panic(err3.Error())
		}
	} else {
		log.Printf("%s already exists \n", clusterRoleBindingName)
	}

	log.Printf("MASTER_URL=%s \n", config.Host)

	// kubectl describe sa default -n kube-system
	sa, _ := clientset.CoreV1().ServiceAccounts("kube-system").Get(context.TODO(), "default",
		metav1.GetOptions{})
	tokenName := sa.Secrets[0].Name
	log.Printf("TOKEN_NAME=%s", tokenName)

	// kubectl get secret default-token-hw9dn -n kube-system -o yaml
	secret, _ := clientset.CoreV1().Secrets("kube-system").Get(context.TODO(), tokenName, metav1.GetOptions{})
	log.Printf("CACRT=%s\n", base64.StdEncoding.EncodeToString(secret.Data["ca.crt"]))
	log.Printf("TOKEN=%s\n", base64.StdEncoding.EncodeToString(secret.Data["token"]))

	list, err4 := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{
		FieldSelector: "spec.unschedulable=false",
	})
	if err4 != nil {
		panic(err4.Error())
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
}
