export default {
  	name: 'notifications',

	/*
	Inject the controller into routes so that it can be accessed
	conveniently by "this.notifier"
	 */
  	initialize: function(container, app) {
    	app.inject('route', 'notifier', 'controller:notifications');
  	}
};
