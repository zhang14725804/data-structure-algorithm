workplace = ~/Desktop/fubao-learning/mysql45
commit_reason = "临时文件"
push:
	cd $(workplace)
	git status
	git add .
	git commit -m $(commit_reason)
	git push origin master
	git status