import Ember from "ember";
import DS from "ember-data";

Ember.$.ajaxPrefilter(function( options, oriOptions, jqXHR ) {
     console.log('setting up auth-token', localStorage.authToken);
     jqXHR.setRequestHeader("auth-token",  localStorage.authToken);
});

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
  /*
  headers: Ember.computed(function(){
  		var token = localStorage.authToken;
      console.log('Authtoken in header set', token);
  		return {"Authtoken": token};
  })
  headers: {
    "auth-token" : localStorage.authToken,
    "Craptoken" : "IAMCRAP"
  }*/
});
