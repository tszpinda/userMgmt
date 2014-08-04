import Ember from 'ember';

/*
The titleize helper converts camelCased attributes into human
readable titles, with the first letter capitalized
 */
export default Ember.Handlebars.makeBoundHelper(function(value) {
  return value.underscore().replace("_", " ").capitalize();
});
