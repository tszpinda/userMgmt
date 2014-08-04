import Ember from 'ember';

/*
The NotificationsController wraps an array of notications that
are shown to the user when various events happen
 */
export default Ember.ArrayController.extend({
	info: function(message) {
    	this.notify(message, 'info');
  	},
  	success: function(message) {
  		this.notify(message, 'success');
  	},
  	error: function(message) {
    	this.notify(message, 'danger');
  	},
  	notify: function(message, type) {
    	var notification = {
      		message: message,
      		type: 'alert-' + type
    	};

    	this.insertAt(0, notification);

    	/*Hide the notication again after 5 seconds */
    	Ember.run.later(this, 'removeObject', notification, 3000);
  	},
  	actions: {
    	remove: function(notification) {
      		this.removeObject(notification);
    	}
  	}
});
