install_node_ubuntu:
	curl -sL https://deb.nodesource.com/setup_15.x | sudo -E bash -
	sudo apt-get install -y nodejs npm


install_node_mac:
	brew install node