package main

var css = `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`