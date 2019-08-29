const express = require('express');
const app = express();

app.get('/api/nodejs', (req, res) => {
	res.send({
		hello: 'world',
		from: 'a service written in nodejs',
		'message sent': ''
	});
});

app.listen('3000', () => console.log('Node.js server running'));