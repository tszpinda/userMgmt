import AuthRoute from './auth';

export default AuthRoute.extend({
	model: function() {
	    /*
	    We don't want to show organizations in the list until they
	    are saved to the server
	     */
	     return this.store.find('tutorial');
	    /*this.store.filter('tutorial', function(record) {
	    	alert('b')
	      return !record.get('isNew');
	    });*/
	}
});
