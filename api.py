from flask import Flask, jsonify, request
from flask_cors import CORS
import random
import string

def random_str(n):
   return ''.join([random.choice(string.ascii_letters + string.digits) for i in range(n)])


app = Flask(__name__)
CORS(app, resources={r"/*": {"origins": "*"}})

@app.route('/garbled', methods = ['POST', 'GET'])
def garbled():
   if request.method == 'POST':
      li = []
      threshold = float(request.form['threshold'])
      l = request.form['contentData']
      for s in l:
          if threshold > random.random():
              li.append(int(random.randrange(0,65535)).to_bytes(2,byteorder='little').decode('shift-jis','replace'))
          elif 0.25 > random.random():
              li.append(random_str(1))
          else:
              li.append(s)

      result = ''.join(li)
      return jsonify(result), 201

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=3000, debug=True)
