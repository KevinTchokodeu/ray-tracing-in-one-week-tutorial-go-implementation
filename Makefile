git-%: 
	git add . 
	git commit -m"$(@:git-%=%)"
	git push 
	clear
	git status
	clear

tag-%:	
	git tag "$(@:tag-%=%)"	
	clear

image-%:
	go run main.go > images/$(@:image-%=%).ppm
	open images/$(@:image-%=%).ppm	
	clear