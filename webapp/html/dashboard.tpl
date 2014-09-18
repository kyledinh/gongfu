    <div><h1>Chirpboard</h1>
        <p class="lead">This is your main screen to get info and to take action.
        </p>

        <h3>Account </h3>       
        <div class="media">
            <a class="pull-left" href="#">
            <img class="media-object" src="img/ronin.jpg" alt="">
            </a>
            <div class="media-body">
            <h4 class="media-heading"><span class="btn btn-xs btn-warning glyphicon glyphicon-pencil" ng-click="editUser({{user.id}})"></span>  {{user.name}}</h4>
            <h5>Id: {{user.id}}</h5>
            <h5>Email: {{user.email}}</h5>            
            </div>
        </div>
        <br/>

        <h3>Chirps <span class="btn btn-xs btn-success glyphicon glyphicon-plus" ng-click="openMod('mod-new-chirp')"></span> </h3>

        <table class="table table-striped">
        	<thead>
        	<tr>
                <th>Date</th>                
        		<th>User</th>
                <th>Type</th>              
                <th>Message</th>                  
        	</tr>
        	</thead>  
        	<tr ng-repeat="row in chirps | orderBy:'id'">
                <td><span class="btn btn-xs btn-warning glyphicon glyphicon-pencil" ng-click="editChirp({{row.id}})"></span>
                    {{row.create_date | date : 'dd MMM yyyy - hh:mm'}}</td>
        		<td>{{row.userid}}</td>  
                <td>{{row.type}}</td>                            
        		<td>{{row.message}}</td>

        	</tr>
        </table>  

        <br/>

    </div>
    <div class="clearfix"></div>
     
    <!-- MODAL -->
    <div class="modal fade" tabindex="-1" id="mod-chirp" aria-hidden="true" style="display: none;">
    <div class="modal-dialog"> 
    <div class="modal-content">

        <div class="modal-header">
            <button type="button" class="close" ng-click="closeMod('mod-chirp')" aria-hidden="true">×</button>
            <h3 id="mod-chirp-header">Chirp ID: {{chirp.id}}</h3>
        </div>
        <div class="modal-body">
 
            <br/>         
            <form role="form">

                <div class="form-group">
                    <label for="chirp-type">Type</label>
                    <input type="text" class="form-control" ng-model="chirp.type" id="chirp-type">
                </div>
                <div class="form-group">
                    <label for="chirp-message">Message</label>
                    <input type="text" class="form-control" ng-model="chirp.message" id="chirp-message">
                </div>                                

                <div class="accordion-heading">
                    <button class="btn btn-warning btn-sm" ng-click="toggleAccordion('collapseOne')">Advanced Options</button>  
                </div><!-- accordion-heading -->

            </form>

        </div><!-- modal-body -->
        <div id="collapseOne" class="modal-body accordion-body collapse border-top">
            <div class="accordion-inner">
                <h5>Advanced Options:</h5> 
                <button class="btn btn-danger" ng-click="deleteChirp({{chirp.id}})">DELETE</button>
            </div>
        </div><!-- collapseOne -->
        <div class="modal-footer">
            <button class="btn btn-primary" ng-click="patchChirp()">Save changes</button>
            <button class="btn" ng-click="closeMod('mod-chirp')">Cancel</button>            
        </div>

    </div><!-- modal-content -->
    </div><!-- modal-dialog -->
    </div><!-- Modal -->      


    <!-- MODAL -->
    <div class="modal fade" tabindex="-1" id="mod-user" aria-hidden="true" style="display: none;">
    <div class="modal-dialog"> 
    <div class="modal-content">

        <div class="modal-header">
            <button type="button" class="close" ng-click="closeMod('mod-user')" aria-hidden="true">×</button>
            <h3 id="mod-chirp-header">Edit Account Info ID: {{user.id}}</h3>
        </div>
        <div class="modal-body">
            <form role="form">
                <div class="form-group">
                    <label for="account-name">Name</label>
                    <input type="text" class="form-control" ng-model="user.name" id="account-name" placeholder="{{account.name}}">
                </div>   
                <div class="form-group">
                    <label for="account-email">Email</label>
                    <input type="text" class="form-control" ng-model="user.email" id="account-email" placeholder="{{account.email}}">
                </div>
            </form>

        </div><!-- modal-body -->
        <div class="modal-footer">
            <button class="btn btn-primary" ng-click="patchUser()">Save changes</button>
            <button class="btn" ng-click="closeMod('mod-user')">Cancel</button>            
        </div>

    </div><!-- modal-content -->
    </div><!-- modal-dialog -->
    </div><!-- Modal -->   


    <!-- MODAL -->
    <div class="modal fade" tabindex="-1" id="mod-new-chirp" aria-hidden="true" style="display: none;">
    <div class="modal-dialog"> 
    <div class="modal-content">

        <div class="modal-header">
            <button type="button" class="close" ng-click="closeMod('mod-new-chirp')" aria-hidden="true">×</button>
            <h3 id="mod-chirp-header">Add A New Chirp</h3>
        </div>
        <div class="modal-body">
            <form role="form">
                <div class="form-group">
                    <label for="account-name">Type</label>
                    <input type="text" class="form-control" ng-model="neu.type" id="neu-type">
                </div>   
                <div class="form-group">
                    <label for="account-email">Message</label>
                    <input type="text" class="form-control" ng-model="neu.message" id="neu-message">
                </div>
            </form>

        </div><!-- modal-body -->
        <div class="modal-footer">
            <button class="btn btn-primary" id="neu-chirp" ng-click="postChirp()">Chirp!</button>
            <button class="btn" ng-click="closeMod('mod-new-chirp')">Cancel</button>            
        </div>

    </div><!-- modal-content -->
    </div><!-- modal-dialog -->
    </div><!-- Modal -->  


    <div id="backdrop" class="fade in"></div>   
