.PHONY: air tailwind clean

air: tailwind
	go build -o ./tmp/main .

tailwind: views/*.go input.css
	./bin/tailwindcss -i input.css -o ./assets/main.css

clean:
	rm -rf tmp
