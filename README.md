#install mongo
wget http://fastdl.mongodb.org/osx/mongodb-osx-x86_64-2.6.4.tgz
tar xzvf mongodb-osx-x86_64-2.6.4.tgz
#mongo start command:
cd bin
mkdir mongodb-data
./mongod --dbpath ../mongodb-data

#install ember-cli
sudo npm install -g ember-cli
#install bower ui dependancy managment
sudo npm install -g bower

#get source clone app
git clone https://github.com/tszpinda/userMgmt.git
cd userMgmt
#change to branch tutorial
git fetch 
git checkout tutorials

#start server service rest api
cd userMgmt/api
go run server.go

#open another termianl and start service ui
npm install
bower install

ember server
#open localhost:4200

#register
#login...
