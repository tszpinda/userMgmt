import Ember from 'ember';

var Router = Ember.Router.extend({
  location: UserMgmtENV.locationType
});

Router.map(function() {
	this.route('signup');
  	this.route('login');  	
  this.route('dashboard');
});

export default Router;
