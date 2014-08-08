import AuthRoute from './../auth';

export default AuthRoute.extend({
	setupController: function(controller, model) {
    	debugger
    	controller.set('model', model);
	}
});
