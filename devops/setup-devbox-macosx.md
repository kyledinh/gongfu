# Setting up a devbox on Mac OSX
Using a VirtualBox to create a "devbox" will provide a virtual server with the same enviroment as the production server. Using Vagrant makes it easy to create and run this VirtualBox. The devbox will run the same OS, libraries and database as production.

## Install on Mac
* VirtualBox - https://www.virtualbox.org/wiki/Downloads  
* Vagrant - http://downloads.vagrantup.com/ 
* Edit `ubuntu-13.10-64-install.sh` and `Vagrantfile` if you want to change the defaults 
* In `devops/`, run `vagrant init` one-time to create the instance. The command uses the Vagrantfile to create the VirtualBox with the Linux OS, then uses `bootstrap.sh` to install and configure Postgres, Node and all libraries need to run the instance. 

## Running the Instance
* `vagrant up` to start the instance/VirtualBox
* `vagrant ssh` to access the instance
* `vagrant halt` to gently stop the instance 
* In the instance, `/home/vagrant/code/` is sym-linked back to the host's `../.` directory. 
* Point your browser to localhost:8000 to get to the instances port :8000

## Command Vagrant Commands
```
vagrant up
vagrant ssh
vagrant suspend | halt | destroy
vagrant box list
vagrant box add {title} {url}
vagrant init {title} 
vagrant reload
```

## GuestAdditions
* On host: `vagrant plugin install vagrant-vbguest`
* On Guest: `/etc/init.d/vboxadd setup`

## References 
* Vagrant Docs - http://docs.vagrantup.com/v2/
* https://wiki.postgresql.org/wiki/PostgreSQL_For_Development_With_Vagrant
* VMs are in `/Users/kyle/VirtualBox VMs`
* images are in `/Users/kyle/.vagrant.d/boxes`
* `<virtualbox-name>/0/virtualbox/box-disk1.vmdk`
