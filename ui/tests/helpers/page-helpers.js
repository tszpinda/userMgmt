import Ember from 'ember';


var pageHelpers = function() {
	//Ember.Test.registerHelper('loginUser', function (app, email, password) {
	//});
	Ember.Test.registerAsyncHelper('loginUser', function (app, email, password) {
		visit("/login");
		fillIn("#email", email);	
 		fillIn("#password", password); 		
 		click("button[type=submit]");
	});

	Ember.Test.registerAsyncHelper('registerUser', function (app, name, email, password) {
		visit("/signup");
		fillIn("#name", name);
 		fillIn("#email", email);
 		fillIn("#password", password);
 		fillIn("#password-confirmation", password);
 		click("button[type=submit]");
	});
}();

export default pageHelpers;



