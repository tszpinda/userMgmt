import Ember from "ember";

var SignupRoute = Ember.Route.extend({
	model: function() {
    	return this.store.createRecord('user');
  	},

  	actions: {
	    signup: function() {
	      var _this = this;

	      this.currentModel.get('errors').clear();

	      this.currentModel.save().then(function() {
	        _this.notifier.success("You have signed up successfully");
	        _this.transitionTo('login');
	      }, function() {
	        return _this.notifier.error("Signup failed");
	      });
    	}
  	}
});

export default SignupRoute;