import Ember from 'ember';


/*
The ApplicationRoute is always active no matter which other routes
are also active. The model for the application route in this app is
the current session, and some of the actions on the route deal with
handling this
 */
export default Ember.Route.extend({

	model: function() {
    	return this.store.find('session', 'current');
  	},
  	
  	setupController: function(controller, session) {
//    	var currentOrganization, organizations;

  //  	App.CSRF_TOKEN = session.get('csrfToken');
    	this.controllerFor('session').set('model', session);

    /*organizations = session.get('user.organizations');

    this.controllerFor('organizations').set('model', organizations);

    currentOrganization = organizations && organizations.get('firstObject');

    if (currentOrganization) {
      this.controllerFor('organization').set('content', currentOrganization);
    }*/
  	},

  	actions: {
    	resetSession: function() {
	    	var _this = this;        
        localStorage.removeItem("authToken");
        this.currentModel.reload().then(function(session){          
          console.log('reset session, token removed: ', localStorage.authToken, session.id);
          _this.currentModel.set('user', null);
          _this.transitionTo('login');          
        });
	      	/*
          this.currentModel.reload().then(function(session) {            
//	      		App.CSRF_TOKEN = session.get('csrfToken');
				      console.log(session);
	        	_this.transitionTo('index');
	      	});*/
    	},

    	logout: function() {
      	  var _this = this;
          this.currentModel.save().then(function() { 
            console.log('logout after save', localStorage.authToken);
        		_this.notifier.success("You have been logged out");
        		_this.send('resetSession');
      		});
    	}
	}
});
