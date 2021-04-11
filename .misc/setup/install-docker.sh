
# Tested on Ubuntu 20.04
# From : https://docs.docker.com/engine/install/ubuntu/
sudo apt-get update

#
sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg \
    lsb-release

#
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | \
     sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

#
echo \
  "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null


# Install
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io


# Add user to docker group
# so that we don't have to prefix docker commands with sudo
sudo usermod -aG docker <username>
newgrp docker # applly the group immediately

# Verify
docker run hello-world