Vagrant.configure("2") do |config|
  config.vm.box = "hashicorp/bionic64"

    # Manager VM configuration
    config.vm.define "hlf" do |hlf|
      hlf.vm.network "forwarded_port", guest: 80, host: 8081, host_ip: "127.0.0.1"
      hlf.vm.network "forwarded_port", guest: 9443, host: 9443
      hlf.vm.network "forwarded_port", guest: 8080, host: 8080
      hlf.vm.network "private_network", ip: "192.168.33.10"

      hlf.vm.synced_folder "D:/raft/HLF", "/home/ubuntu"

      hlf.vm.provision "shell", inline: <<-SCRIPT
        # Rename the vagrant user to hlf
        sudo usermod -l ubuntu vagrant
        sudo groupmod -n ubuntu vagrant
        sudo usermod -d /home/ubuntu -m ubuntu

        # Update sudoers file
        sudo sed -i 's/vagrant ALL=(ALL:ALL) NOPASSWD:ALL/ubuntu ALL=(ALL:ALL) NOPASSWD:ALL/' /etc/sudoers

        # Update profile and bashrc
        echo "cd /vagrant" >> /home/ubuntu/.profile
        echo "cd /vagrant" >> /home/ubuntu/.bashrc
        echo "All good!!"
      SCRIPT
    end

#   # Manager VM configuration
#   config.vm.define "manager" do |manager|
#     manager.vm.network "forwarded_port", guest: 80, host: 8081, host_ip: "127.0.0.1"
#     manager.vm.network "forwarded_port", guest: 9443, host: 9443
#     manager.vm.network "forwarded_port", guest: 8080, host: 8080
#     manager.vm.network "private_network", ip: "192.168.33.10"
#
#     manager.vm.synced_folder "D:/raft/HLF", "/home/ubuntu/hlf-docker-swarm"
#
#     manager.vm.provision "shell", inline: <<-SCRIPT
#       # Rename the vagrant user to manager
#       sudo usermod -l ubuntu vagrant
#       sudo groupmod -n ubuntu vagrant
#       sudo usermod -d /home/ubuntu -m ubuntu
#
#       # Update sudoers file
#       sudo sed -i 's/vagrant ALL=(ALL:ALL) NOPASSWD:ALL/ubuntu ALL=(ALL:ALL) NOPASSWD:ALL/' /etc/sudoers
#
#       # Update profile and bashrc
#       echo "cd /vagrant" >> /home/ubuntu/.profile
#       echo "cd /vagrant" >> /home/ubuntu/.bashrc
#       echo "All good!!"
#     SCRIPT
#   end
#
#   # Worker1 VM configuration
#   config.vm.define "worker1" do |worker1|
#     worker1.vm.network "forwarded_port", guest: 80, host: 8082, host_ip: "127.0.0.1"
#     worker1.vm.network "private_network", ip: "192.168.33.11"
#
#     worker1.vm.synced_folder "D:/raft/HLF", "/home/ubuntu/hlf-docker-swarm"
#
#         worker1.vm.provision "shell", inline: <<-SCRIPT
#           # Rename the vagrant user to manager
#           sudo usermod -l ubuntu vagrant
#           sudo groupmod -n ubuntu vagrant
#           sudo usermod -d /home/ubuntu -m ubuntu
#
#           # Update sudoers file
#           sudo sed -i 's/vagrant ALL=(ALL:ALL) NOPASSWD:ALL/ubuntu ALL=(ALL:ALL) NOPASSWD:ALL/' /etc/sudoers
#
#           # Update profile and bashrc
#           echo "cd /vagrant" >> /home/ubuntu/.profile
#           echo "cd /vagrant" >> /home/ubuntu/.bashrc
#           echo "All good!!"
#         SCRIPT
#   end
#
#   # Worker2 VM configuration
#   config.vm.define "worker2" do |worker2|
#     worker2.vm.network "forwarded_port", guest: 80, host: 8083, host_ip: "127.0.0.1"
#     worker2.vm.network "private_network", ip: "192.168.33.12"
#
#     worker2.vm.synced_folder "D:/raft/HLF", "/home/ubuntu/hlf-docker-swarm"
#
#             worker2.vm.provision "shell", inline: <<-SCRIPT
#               # Rename the vagrant user to manager
#               sudo usermod -l ubuntu vagrant
#               sudo groupmod -n ubuntu vagrant
#               sudo usermod -d /home/ubuntu -m ubuntu
#
#               # Update sudoers file
#               sudo sed -i 's/vagrant ALL=(ALL:ALL) NOPASSWD:ALL/ubuntu ALL=(ALL:ALL) NOPASSWD:ALL/' /etc/sudoers
#
#               # Update profile and bashrc
#               echo "cd /vagrant" >> /home/ubuntu/.profile
#               echo "cd /vagrant" >> /home/ubuntu/.bashrc
#               echo "All good!!"
#             SCRIPT
#   end
end
