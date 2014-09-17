      <div class="margin-60">
        <h1>Sign Up Form</h1>
        <p class="lead">Use the form below to create an account.<br> 
         </p>

        <div class="margin-20 col-lg-6 col-md-6">
          <form role="form" action="/signup" method="POST" >
            <div class="form-group">
              <label for="inputName">Name</label>
              <input type="text" name="name" class="form-control" id="inputName" placeholder="Enter name">
            </div>            
            <div class="form-group">
              <label for="inputEmail">Email address</label>
              <input type="email" name="email" class="form-control" id="inputEmail" placeholder="Enter email">
            </div>
            <div class="form-group">
              <label for="inputPassword1">Password</label>
              <input type="password" name="password" class="form-control" id="inputPassword1" placeholder="Password">
            </div>
            <div class="form-group">
              <label for="inputPassword2">Retype Password</label>
              <input type="password" name="password2" class="form-control" id="inputPassword2" placeholder="Password">
            </div>            
            <button type="submit" id="btn_create_account" class="btn btn-success">Create Account</button>

            <a href="/webapp/app.html#/" class="btn btn-info">Already Have An Account</a>
          </form>
        </div>         
      </div>
