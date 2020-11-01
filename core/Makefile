.PHONY: templates

run:
	go run main.go

gen:
	$$HOME/go/bin/qtc -skipLineComments 

clean:
	find . -name *.qtpl.go | tee /dev/tty | xargs rm -f 
