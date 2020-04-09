import sqlite3

if __name__ == '__main__':
    conn = sqlite3.connect('data.db')
    c = conn.cursor()
    rows = c.execute("SELECT * from t WHERE m LIKE '{}'".format('c4ca4238a0b923820dcc509a6f75849b%'))
    for row in rows:
        print('md5: {m}, value: {v}, type: {t}'.format(m=row[1], v=row[0], t=row[2]), end='\n')

    conn.close()