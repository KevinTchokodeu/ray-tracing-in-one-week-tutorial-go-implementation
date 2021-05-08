git-%: 
	git add . 
	git commit -m"$(@:git-%=%)"
	git push 
	clear
	git status
	clear			