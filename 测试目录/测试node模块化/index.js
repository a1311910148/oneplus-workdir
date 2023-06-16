const mod = require('./mod');
const express = require('express');

// ls.stdout.on('data', (res) => console.log(res.toString()))
let app = mod.server(3000, (req, res) => res.send('hello world'))


