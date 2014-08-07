import Ember from 'ember';
import startApp from '../helpers/start-app';

var App;

var name = "John";
var email;
var password = "Password1";

module('Acceptance: RegisterLogin', {
  setup: function() {
  	window.localStorage.clear();
    App = startApp();
	email = "john@" + new Date().getTime() + ".com";
  },
  teardown: function() {
    Ember.run(App, 'destroy');
  }
});


test('register', function() {
	expect(2);

	App.testHelpers.registerUser(name, email, password); 
	andThen(function(){		
   		equal(currentPath(), 'login');
	});
	App.testHelpers.loginUser(email, password); 
	andThen(function(){
		equal(currentPath(), 'dashboard');
	});
});

test('authenticated only', function() {
	visit('/dashboard');
	andThen(function(){
		equal(currentPath(), 'login');
	});
});
