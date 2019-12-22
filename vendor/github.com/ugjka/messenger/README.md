# Messenger
Simple broadcasting mechanism using go channels

[![Build Status](https://travis-ci.org/ugjka/messenger.svg?branch=master)](https://travis-ci.org/ugjka/messenger)
[![codecov](https://codecov.io/gh/ugjka/messenger/branch/master/graph/badge.svg)](https://codecov.io/gh/ugjka/messenger)
[![GoDoc](https://godoc.org/github.com/ugjka/messenger?status.svg)](https://godoc.org/github.com/ugjka/messenger)
[![Go Report Card](https://goreportcard.com/badge/github.com/ugjka/messenger)](https://goreportcard.com/report/github.com/ugjka/messenger)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Donate](https://dl.ugjka.net/Donate-PayPal-green.svg)](https://www.paypal.me/ugjka)

## What is this

This is a simple broadcasting mechanism where gouroutines can subscribe and unsubscribe to recieve messages from the Messenger instance.

Channels can be buffered to reduce blocking and messages can be dropped if channel's buffer is full. These are configurable options.

## Why

Because i needed this for some silly stuff and ofcourse, why not!