from bottle import Bottle
import psycopg2

product = Bottle()

@product.get('/api/product/search')
def searchProduct():
    pass

@product.delete('/api/product/{id}/remove-favorite')
def deleteProduct():
    pass

@product.post('/api/product/{id}/add-favorite')
def addProduct():
    conn = psycopg2.connect("user='pr0n00gler' password='pass' host='localhost' dbname='hack'")
    db = conn.cursor()
    querry = "insert into product(id,company_id,category_id, name) values ({0}, {1}, {2}, '{3}')".format(id, 1,2,'молоко')
    db.execute(querry)
    db.close()
    conn.close()
    pass

@product.get('/api/product/{id}')
def getProduct():
    pass