import Ember from "ember";
import DS from "ember-data";

export default DS.RESTAdapter.extend({
	host: 'http://localhost:3000',
	ajaxError: function(jqXHR) {
    	var error = this._super(jqXHR);

        if (jqXHR && jqXHR.status === 422) {
        	var jsonErrors = Ember.$.parseJSON(jqXHR.responseText)["errors"];
	  		return new DS.InvalidError(jsonErrors);
        } else {
        	return error;
        }
    },
    headers: Ember.computed(function(){
  		var token = localStorage.authToken;
  		return {"AuthToken": token};
  	})
});
