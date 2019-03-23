from bottle import Bottle
from order import order
#from products import product

rootApp = Bottle()
@rootApp.route('/')
def rootIndex():
    return 'Application Suite Home Page'

if __name__ == '__main__':
    rootApp.merge(order)
    rootApp.run(debug=True, port=8080)