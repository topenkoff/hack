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
    bodyRequest = json.load(request.body)
    conn = psycopg2.connect("user='pr0n00gler' password='pass' host='localhost' dbname='hack'")
    db = conn.cursor()
    query = "insert into order values ('{0}','{1}','{2}','{3}')".format(bodyRequest[0],bodyRequest[1],bodyRequest[2],bodyRequest[3])
    db.execute(query)
    db.close()
    conn.close()
    body = {}
    body.update(query)
    return HTTPResponse(status=201, body=body)


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
    db_order = getValue('select * from orders where id = {id}'.format(id))
    _user_id = db_order[0]
    _number = db_order[1]
    _payment_method = db_order[2]
    body = {"id":id, "user_id":_user_id, "number":_number,"payment_method":_payment_method}
    return HTTPResponse(status=200, body=body)

@order.put('/api/order/<id>')
def updateOrder():
    pass


@order.get('/api/order/<id>')
def getOrder(id):
    db_order = getValue('select * from orders where id = {id}'.format(id))
    _user_id = db_order[0]
    _number = db_order[1]
    _payment_method = db_order[2]
    body = {"id":id, "user_id":_user_id, "number":_number,"payment_method":_payment_method}
    body = {}
    return HTTPResponse(status=200, body=body)


@order.get('/api/order')
def getOrders():
    pass
    

