import AuthRoute from './../auth';


export default AuthRoute.extend({
	model: function() {
		debugger
		return this.store.find('tutorial', '53e3e3fc421aa977e0000004')
    	//return this.store.find('session', 'current');
  	},
    setupController: function(controller, model) {
    	debugger
    	controller.set('model', model);
	}	
});
