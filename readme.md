# Gongfu

Gongfu is a mashup of Go and (ng) Angular used together to build a website. This template will get you started to build a Smart Angular Client and a RESTful API Go server. This template uses Vagrant/VirtualBox to build a virtual environment and NodeJS/Grunt to do task building.

This is the source code to a presentation I gave at the SF Golang meetup in Sept. Here's the link to it: http://www.hakkalabs.co/articles/angularjs-go-lean-combination#!

## App Stack
* Ubuntu 13.10 64 bit (http://www.ubuntu.com/)
* Go 1.3 (http://golang.org/)
* Gin-Gonic/Gin (https://github.com/gin-gonic/gin) 
* Angular (https://angularjs.org/)
* PostgreSQL (http://www.postgresql.org/)
* Bootstrap (http://getbootstrap.com/)

* Nodejs (http://nodejs.org/)
* Gruntjs (http://gruntjs.com/)

## Dev Stack
* Vagrant
	* `vagrant up | halt | destroy | ssh`
* Grunt  
	* build
	* package
	* test
	* deploy

## Gongfu setup
Step 1: Build a Vagrant instance
Step 2: Configure instance
Step 3: Compile Gongfu app
Step 4: Launch Gongfu app

### Step 1: Build a Vagrant/Virtualbox instance

On your host/dev machine you will need to install

* Vagrant (https://www.vagrantup.com/downloads.html)
* VirtualBox (https://www.virtualbox.org/wiki/Downloads)

The `devops` directory will be will home directory for Vagrant and your admin tools and scripts. In the directory you will start the instance with:

* `vagrant up`

The Vagrantfile
```
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  # Ubuntu-13.10
  config.vm.box = "Ubuntu-13.10"
  config.vm.box_url = "http://opscode-vm-bento.s3.amazonaws.com/vagrant/virtualbox/opscode_ubuntu-13.10_chef-provisionerless.box"
  config.vm.synced_folder "../.", "/home/vagrant/code"
  config.ssh.forward_x11 = true
  config.vm.network "forwarded_port", guest: 8000, host: 8000, auto_correct: true
  config.vm.network "forwarded_port", guest: 3000, host: 3000, auto_correct: true
  config.vm.network "forwarded_port", guest: 5432, host: 5432, auto_correct: true
  config.vm.network "private_network", ip: "192.168.10.10"
  config.vm.provider :virtualbox do |virtualbox|
    virtualbox.customize ["modifyvm", :id, "--name", "gongfu"] # Name the INSTANCE (app), make unique
    virtualbox.customize ["modifyvm", :id, "--memory", "2048"]
  end
  config.vm.provision :shell, :path => "ubuntu-13.10-64-install.sh"
end

```


### Step2: Configure instance

This application and installation was designed to driven by one config file `cfg.json`. The `ubuntu-13.10-64-install.sh` script will configure your Postgres database with these settings. The Grunt build tool will also use this cfg file for its settings. As well as the Gongfu app. I would modify all the "gongfu" to your app's name and of course change the "db_password".

cfg.json
```
{
   "appname" : "gongfu",
   "http_port" : "",
   "server" : "localhost",
   "db_server" : "localhost",
   "db_name" : "gongfu",
   "db_user" : "gongfu",
   "db_password" : "kungfupanda"
}
```

* `vagrant ssh`  will log you into the instance
* In the instance you will see that `/home/vagrant/code` is a sym-link to your host's repo
* So if you, `cd ~/code/devops`, you will be in the devops dir in your instance and you can run
* `ubuntu-13.10-64-install.sh` to configure the instance
* Use `sys-check.sh` to see your status and version of software installed
* Use `go-packages.sh` to install the Go packages


### Step 3: Compile Gongfu app

* cd back to `cd ~/code` to get back to the source root directory
* `sudo npm install` will install the NodeJS packages for the Grunt build tool
* You can use `devops/sys-check.sh` to check your installation
* From here you want to look in the `Gruntfile.js` for the tasks available

partial Gruntfile.js
```
  grunt.registerTask('dev', ['files', 'shell:gotest', 'go:build:myapp', 'shell:run']);
  grunt.registerTask('production', ['files', 'shell:gotest', 'go:build:myapp']);

  // database tasks
  grunt.registerTask('pgusers', ['shell:psqlusers']);
  grunt.registerTask('pgcreate', ['shell:psqlcreate']);
  grunt.registerTask('pgdrop', ['shell:psqldrop']);
  grunt.registerTask('pgdump', ['shell:psqldump']);
  grunt.registerTask('pglist', ['shell:psqllist']);
  grunt.registerTask('pgrebuild', ['shell:psqldrop', 'shell:psqlcreate']);
  grunt.registerTask('pgrestore', ['shell:psqlrestore']);

  // default build is dev
  grunt.registerTask('build', ['dev']);
  grunt.registerTask('test', ['e2e', 'unit']);
  grunt.registerTask('run', ['build', 'shell:run']);
  grunt.registerTask('got', ['shell:gotest']);
  grunt.registerTask('pack', ['compress:main']);
```
* You can use `grunt build` that will launch a "dev" task, that a macro for "grunt.registerTask('dev', ['files', 'shell:gotest', 'go:build:myapp', 'shell:run']);". This one task will concat files, run the go unit tests, build the go executable and run that executable.


### Step 4: Running the Gongfu app

* You can also just run the app by using the executable `sh gongfuServer`
* From your host machine you can reach the Gongfu instance at `192.168.10.10`, this address can be changed in your Vagrantfile.
* When you want to terminate the instance, just exit back out to the host and run `vagrant halt` in the "devops" directory.



