import Ember from 'ember';

/*
The AuthenticatedRoute is a parent class for routes that require
login. It will redirect to the login page if there is no current
user. This is only for UI purposes and does not provide any
security, the server must authenticate all endpoints regardless
 */
export default Ember.Route.extend({
	beforeModel: function() {
		var user = this.modelFor('application').get('user');
		var loggedIn = user && user.get('name') !== "";
		console.log("loggedIn", loggedIn);
		if(loggedIn){
			this.transitionTo('dashboard');			
	    }
  	}
});