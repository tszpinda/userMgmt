import AuthRoute from '../auth';
//import Ember from 'ember';

export default AuthRoute.extend({
	
	model: function(args) {
		console.log("route:tutorial/view", args.tutorial_id);
	},

	setupController:function(controller,model){
    console.log("in setupController hook for login route");
  }
})