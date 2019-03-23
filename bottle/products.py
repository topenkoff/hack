from bottle import Bottle

product = Bottle()

@product.get('/api/product/search')
def searchProduct():
    pass

@product.delete('/api/product/{id}/remove-favorite')
def deleteProduct():
    pass

@product.post('/api/product/{id}/add-favorite')
def addProduct():
    pass

@product.get('/api/product/{id}')
def getProduct():
    pass