import Ember from "ember";

export default Ember.Handlebars.makeBoundHelper(function(value) {
    console.debug('hbs template:', value);
});

