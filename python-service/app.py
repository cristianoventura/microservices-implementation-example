from flask import Flask, jsonify

app = Flask(__name__)

@app.route('/', methods=['GET'])
def get():
	return jsonify({
		'hello': 'world',
		'from': 'a service written in python',
		'message sent': ''
	}), 200

if __name__ == '__main__':
	app.run(host='0.0.0.0', port=3000)