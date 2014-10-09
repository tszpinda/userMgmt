import AuthRoute from '../auth';

export default AuthRoute.extend({
	
	model: function(args) {
        //debugger
		console.log("route:view/steps", args.tutorial_id);
		
		var viewCntrl = this.controllerFor('tutorial');
		var model = viewCntrl.get('model');
		//debugger
		return model.get('steps');
		//return this.store.find('tutorial', args.tutorial_id);
	},

	setupController: function(controller, model) {
		console.log(model)
	}

})