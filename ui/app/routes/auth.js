import Ember from 'ember';

/*
The AuthenticatedRoute is a parent class for routes that require
login. It will redirect to the login page if there is no current
user. This is only for UI purposes and does not provide any
security, the server must authenticate all endpoints regardless
 */
export default Ember.Route.extend({
	beforeModel: function(transition) {
		//debugger
		var requiresAuth = localStorage.authToken === undefined;
		if (!requiresAuth) {
			//verify actully if the token & user are valid
			var session = this.modelFor('application');
			var user = session.get('user');
			if(!user){
				requiresAuth = true;
			}
			console.log('logged in', user.get('name'));
		}
    	if (requiresAuth) {
    		//localStorage.removeItem('authToken');
    		console.log('requiresAuth');
	   		/*   
	      The saved transition is retried after login, so the user can
	      directly be delivered to the original page they requested
	      after they have authenticated
	       */
	      this.controllerFor('session').set('savedTransition', transition);
	      this.transitionTo('login');
	    }else{
	    	console.log('has user');
	    }
  	}
});