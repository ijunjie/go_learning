import sqlite3

import hashlib


# from guppy import hpy

class Record:
    def __init__(self, v, t):
        self.v = str(v)
        self.md5 = hashlib.md5(self.v.encode('utf-8')).hexdigest()
        self.t = t

    def as_tuple(self):
        return (self.v, self.md5, self.t)


TIDS = [x for x in range(1, 100)]
PIDS = [100, 167, 180, 181, 182, 183, 184, 185, 186, 187, 188, 189, 190, 1009, 1010]
GIDS = [x for x in range(1, 100)]

if __name__ == '__main__':
    # h = hpy()
    # heap = h.heap()

    conn = sqlite3.connect('data.db')
    c = conn.cursor()
    c.execute('''CREATE TABLE IF NOT EXISTS t
           (v CHAR(20) NOT NULL,
           m CHAR(32) NOT NULL,
           t CHAR(1) NOT NULL);''')

    insert_tmpl = 'INSERT INTO t VALUES (?,?,?)'
    for tid in TIDS:
        t1 = Record(tid, 't').as_tuple()
        c.execute(insert_tmpl, t1)
        for pid in PIDS:
            t2 = Record(str(tid) + str(pid), 'p').as_tuple()
            c.execute(insert_tmpl, t2)
            for gid in GIDS:
                t3 = Record('_'.join([str(tid), str(pid), str(gid)]), 'g').as_tuple()
                c.execute(insert_tmpl, t3)

    conn.commit()
    conn.close()

    # print(heap)
