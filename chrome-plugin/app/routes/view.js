import AuthRoute from './auth';
//import Ember from 'ember';

export default AuthRoute.extend({
	
	model: function(args) {
		console.log("route:view", args.tutorial_id);
		
		var viewCntrl = this.controllerFor('tutorial');
		var model = viewCntrl.get('model');
		return model;

		console.log("route:view", args.tutorial_id);
		console.log("route:view", args.view_id);
		 return this.store.find('tutorial', args.tutorial_id);
	},
	actions: {
		saveStep:function() {
          console.log('saving step');
        },
        rollbackStep:function() {
          console.log('rollback step');
        }
	}

})