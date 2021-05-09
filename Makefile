git-% tag-%: 
	git add . 
	git commit -m"$(@:git-%=%)"
	git tag "$(@:tag-%=%)"

	clear
	git status
	clear			