import Ember from 'ember';
import startApp from '../helpers/start-app';

var App;

var name = "John";
var password = "Password1";
var email, tutorialName;

module('Acceptance: Tutorial-CURD', {
  setup: function() {
  	window.localStorage.clear();
    App = startApp();
    var unique = new Date().getTime();
	email = "john@" + unique + ".com";
	tutorialName = "create-job-tutorial-" + unique;
	App.testHelpers.registerUser(name, email, password); 
	App.testHelpers.loginUser(email, password); 
  },
  teardown: function() {
    Ember.run(App, 'destroy');
  }
});

test('tutorials list', function() {
	expect(2);
	visit('/tutorials');
	andThen(function(){
		equal(currentPath(), 'tutorials.index');
	});
	click(".js-add-tutorial");
	andThen(function(){
		equal(currentPath(), 'tutorials.new');			
	});
	fillIn('#name', tutorialName);
	fillIn('#domain', "localhost");
	fillIn('#page', "login");
	click('#save');
	//andThen(function(){
	//	equal(currentPath(), 'tutorials/edit');
	//});
});

