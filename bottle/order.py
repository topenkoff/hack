from bottle import Bottle, HTTPResponse, request
import json
import psycopg2

def getValue(query):
    connection_db = psycopg2.connect("user='pr0n00gler' password='pass' host='localhost' dbname='hack'")
    db = connection_db.cursor()
    db.execute(query)
    result = db.fetchall()
    db.close()
    connection_db.close()
    return result[0]

order = Bottle()

@order.post('/api/order')
def createOrder():
    pass


@order.post('/api/order/<id>/i-here')
def clientHere(id):
    return HTTPResponse(status=200, body='')


@order.get('/api/order/get-payment-method')
def getPaymentMethod():
    body = json.dumps(["cash", "e_money"])
    return HTTPResponse(status=200, body=body)


@order.delete('/api/order/<id>')
def deleteOrder():
    conn = psycopg2.connect("user='pr0n00gler' password='pass' host='localhost' dbname='hack'")
    db = conn.cursor()
    querry = "delete from orders where id = {0}".format(id)
    db.execute(querry)
    db.close()
    conn.close()
    body = {}
    return HTTPResponse(status=200, body=body)

@order.put('/api/order/<id>')
def updateOrder():
    pass


@order.get('/api/order/<id>')
def getOrder(id):
    return id


@order.get('/api/order')
def getOrders():
    pass
    

