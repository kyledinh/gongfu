'use strict';

/* jasmine specs for controllers go here */

describe('Algorithms Sorting Functions', function() {
 
   describe('Selection Sort', function(){
      it('should sort this array', function() {
         var arr = ["v","a","d","e","r"];
         var ss = selectionSort(arr);
         var js = arr.sort(); 
         var n = ss.length;

         for (var i=0; i < n; i++) {
            expect(ss[i]).toBe(js[i]);
         } 

      });
   });

   describe('Insertion Sort', function(){
      it('should sort this array', function() {
         var arr = ["v","a","d","e","r"];
         var is = insertionSort(arr);
         var js = arr.sort(); 
         var n = is.length;

         for (var i=0; i < n; i++) {
            expect(is[i]).toBe(js[i]);
         } 

         expect(is[0]).toBe("a");
         expect(is[1]).toBe("d");
         expect(is[4]).toBe("v");
      });
   });

   describe('Quick Sort', function(){
      it('should sort this array', function() {
         var arr = ["v","a","d","e","r"];
         var qs = quickSort(arr);
         var js = arr.sort(); 
         var n = qs.length;

         for (var i=0; i < n; i++) {
            expect(qs[i]).toBe(js[i]);
         } 
         
         expect(qs[0]).toBe("a");
         expect(qs[1]).toBe("d");
         expect(qs[4]).toBe("v");

      });
   });

});
