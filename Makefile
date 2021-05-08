git-%: 
	git add . 
	git commit -m"$(@:git-%=%)"
	git push 
	git clear
	git status			