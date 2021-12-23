workplace = ~/Desktop/fubao-learning/mysql45
commit_reason = "删除链表的第n个节点"
push:
	cd $(workplace)
	git status
	git add .
	git commit -m $(commit_reason)
	git push origin master
	git status