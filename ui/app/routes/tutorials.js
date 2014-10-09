import AuthRoute from './auth';

export default AuthRoute.extend({
	model: function() {
		console.log("route:tutorials");
	    /*
	    We don't want to show organizations in the list until they
	    are saved to the server
	     */
	    //return this.store.find('tutorial');
	    /*return this.store.filter('tutorial', function(record) {
            console.log('filter', record.get('isNew'));
	      return !record.get('isNew');
	    });*/
        return this.store.filter('tutorial', {isNew: false}, function(post) {
            console.log(post.get('isNew'))
            return !post.get('isNew');
        })
	}
});
