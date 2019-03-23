from bottle import Bottle, HTTPResponse, request
import json
import psycopg2


order = Bottle()


@order.post('/api/order')
def createOrder():
    
    #headerAuth = request.get_header('Authorization')
    #if headerAuth == 'Bearer MaX_kMWYU8i3wM9P9ZCs6ERkIbCRY0kV':
    #body = {"id" = _id, "user_id" = _user_id, "number" = _number, "amount" = _amount, "payment_methon" = _payment_method, "user" = {"id" = __id, "username" = __username, "email" = __email, "phone" = __phone, "created_at" = __created_at, "updated_at" = __updated_at}"company" = _company}
    #return HTTPResponse(body=body, status=code)


@order.post('/api/order/<idint>/i-here')
def clientHere(id):
    pass


@order.get('/api/order/get-payment-method')
def getPaymentMethod():
    pass


@order.delete('/api/order/<id>')
def deleteOrder():
    conn = psycopg2.connect("user='pr0n00gler' password='pass' host='localhost' dbname='hack'")
    db = conn.cursor()
    querry = "delete from orders where id = {0}".format(id)
    db.execute(querry)
    db.close()
    conn.close()
    return HTTPResponse(status=200, body='')

@order.put('/api/order/<id>')
def updateOrder():
    pass


@order.get('/api/order/<id>')
def getOrder(id):
    return id


@order.get('/api/order')
def getOrders():
    headerAuth = request.get_header('Authorization')
    bodyRequest = json.load(request.body)
    conn = psycopg2.connect("user='pr0n00gler' password='pass' host='localhost' dbname='hack'")
    cursor = conn.cursor()
    

