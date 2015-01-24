module.exports = function(grunt) {
  grunt.initConfig({

    cfg: grunt.file.readJSON('cfg.json'),

    clean: ['build/*', 'dist/*', 'webapp/dist', 'webapp/js/*'],

    autoprefixer: {
      options: {
        //browsers: configBridge.config.autoprefixerBrowsers
      },
      core: {
        options: {
          map: true
        },
        src: 'dist/css/<%= cfg.bootstrap %>.css'
      },
      theme: {
        options: {
          map: true
        },
        src: 'dist/css/<%= cfg.bootstrap %>-theme.css'
      }
    },

    compress: {
      main: {
        options: {
          archive: 'build/gongfu-<%= cfg.appversion %>.zip'
        },
        files: [
          { expand: true, dest: 'bash', src: '*.sh', cwd: 'devops/' },
          { expand: true, src: 'certs/**' },
          { expand: true, src: 'webapp/**' },
          { expand: true, src: 'cfg.json' },
          { expand: true, src: 'gongfu' }
        ]
      }
    },

    concat: {
      options: {
        separator: ';',
        banner: '',
        stripBanners: false
      },
      appjs: {
        dest: 'webapp/js/<%= cfg.appname %>-app.js', 
        src: [
          'src/angular/app.js',
          'src/angular/ctrl-*.js',
          'src/angular/directives.js',
          'src/angular/factories.js',
          'src/angular/filters.js',          
          'src/angular/util.js'                   
        ]
      },
      bootstrap: {
        src: [
          'src/js/transition.js',
          'src/js/alert.js',
          'src/js/button.js',
          'src/js/carousel.js',
          'src/js/collapse.js',
          'src/js/dropdown.js',
          'src/js/modal.js',
          'src/js/tooltip.js',
          'src/js/popover.js',
          'src/js/scrollspy.js',
          'src/js/tab.js',
          'src/js/affix.js',
          'src/js/swiper-slide-1.9.3.js',
          'src/js/jquery.easing.js'
        ],
        dest: 'dist/js/<%= cfg.bootstrap %>.js'
      },
      spec: {
        dest: 'webapp/js/app.js',
        src: [
          'webapp/js/app.js',
          'src/javascript/cfg-spec.js'
        ]
      }
    }, 

    copy: {
      fonts: {
        src: 'src/fonts/*',
        dest: 'dist/'
      },
      dist_webapp_css: {
        src: 'dist/css/*',
        dest: 'webapp/css/',
        flatten: true,
        expand: true
      },
      dist_webapp_js: {
        src: 'dist/js/*',
        dest: 'webapp/js/',
        flatten: true,
        expand: true
      }      
    },

    cssmin: {
      options: {
        compatibility: 'ie8',
        keepSpecialComments: '*',
        noAdvanced: true
      },
      minifyCore: {
        src: 'dist/css/<%= cfg.bootstrap %>.css',
        dest: 'dist/css/<%= cfg.bootstrap %>.min.css'
      },
      minifyTheme: {
        src: 'dist/css/<%= cfg.bootstrap %>-theme.css',
        dest: 'dist/css/<%= cfg.bootstrap %>-theme.min.css'
      }
    },


    go: {
      options: {
	       GOPATH: ["./", "/home/vagrant/code"]
      },
      myapp: {
        root: './', 
        output: '<%= cfg.appname %>Server' ,
        run_files: ['app.go']
      }
    },

    jshint: {
      src: ['Gruntfile.js', 'src/angular/*.js', 'src/javascript/*.js'],
      options: {
        curly: true,
        eqeqeq: true,
        immed: true,
        latedef: true,
        newcap: true,
        noarg: true,
        sub: true,
        undef: true,
        boss: true,
        eqnull: true,
        browser: true,
        globals: {
          require: true,
          define: true,
          requirejs: true,
          describe: true,
          expect: true,
          it: true
        }
      }
    }, 

    karma: {
      unit: { configFile: 'test/karma.conf.js' } 
    },

    less: {
      compileCore: {
        options: {
          strictMath: true,
          sourceMap: true,
          outputSourceFiles: true,
          sourceMapURL: '<%= cfg.appname %>.css.map',
          sourceMapFilename: 'dist/css/<%= cfg.bootstrap %>.css.map'
        },
        src: 'src/less/bootstrap.less',
        dest: 'dist/css/<%= cfg.bootstrap %>.css'
      },
      compileTheme: {
        options: {
          strictMath: true,
          sourceMap: true,
          outputSourceFiles: true,
          sourceMapURL: '<%= cfg.bootstrap %>-theme.css.map',
          sourceMapFilename: 'dist/css/<%= cfg.bootstrap %>-theme.css.map'
        },
        src: 'src/less/theme.less',
        dest: 'dist/css/<%= cfg.bootstrap %>-theme.css'
      }
    },

    protractor: {
      options: {
        keepAlive: true,
        configFile: 'test/protractor.conf.js',
        args: {
          seleniumServerJar: 'node_modules/protractor/selenium/selenium-server-standalone-2.42.2.jar',
          chromeDriver: 'node_modules/protractor/selenium/chromedriver'
        }
      },
      run: {},
      localhost: {
        options: { configFile: 'test/localhost.conf.js', }
      },
      staging: {
        options: { configFile: 'test/staging.conf.js', }
      }
    },

    shell: {
      gotest: {
        command: [
          'go test ./src/go/gongfu/...',
          'echo "hello"'
        ].join('&&')
      },
      run: {
        command: 'sudo ./<%= cfg.appname %>Server',  
        options: {
          execOptions: {
            cwd: './'
          }
        }
      },
      psqlusers: {
        command: 'psql -h localhost -U <%= cfg.db_user %> -d <%= cfg.db_name %> -c "select * from users" ',  
      },
      psqlcreate: {
        command: 'psql -h localhost -U <%= cfg.db_user %> -d <%= cfg.db_name %> -f devops/sql/create_tables.sql',  
      },
      psqldrop: {
        command: 'psql -h localhost -U <%= cfg.db_user %> -d <%= cfg.db_name %> -f devops/sql/drop_tables.sql',  
      },
      psqldump: {
        command: 'pg_dump -h localhost -U <%= cfg.db_user %> <%= cfg.db_name %> -f devops/sql/<%= cfg.db_name %>.dump.sql',  
      },
      psqllist: {
        command: 'psql -h localhost -U <%= cfg.db_user %> -d <%= cfg.db_name %> -f devops/sql/list_tables.sql',  
      }, 
      psqlrestore: {
        command: 'psql -h localhost -U <%= cfg.db_user %> -d <%= cfg.db_name %> -f devops/sql/<%= cfg.db_name %>.dump.sql',  
      },
    },

    uglify: {
      options: {
        sourceMap: true,
        mangle: false,
        beautify: true,
        sourceMapName: function (filePath) {
          return filePath + '.map';
        }
      },
      layout: {
        files: {
          'webapp/js/app.js': [
            'webapp/js/app.js'
          ]
        }
      }
    },

    watch: {
      files: '<%= jshint.src %>',
      tasks: ['jshint', 'karma']
    }
  });

  // Load JSHint task
  grunt.loadNpmTasks('grunt-autoprefixer');
  grunt.loadNpmTasks('grunt-contrib-clean');
  grunt.loadNpmTasks('grunt-contrib-compress');
  grunt.loadNpmTasks('grunt-contrib-concat');
  grunt.loadNpmTasks('grunt-contrib-copy');  
  grunt.loadNpmTasks('grunt-contrib-cssmin');   
  grunt.loadNpmTasks('grunt-contrib-jshint');
  grunt.loadNpmTasks('grunt-contrib-less');    
  grunt.loadNpmTasks('grunt-contrib-uglify');  
  grunt.loadNpmTasks('grunt-go');
  grunt.loadNpmTasks('grunt-karma');
  grunt.loadNpmTasks('grunt-protractor-runner');
  grunt.loadNpmTasks('grunt-shell');


  // Default task.
  grunt.registerTask('default', ['build','unit']);
  grunt.registerTask('e2e', ['protractor:run']);
  grunt.registerTask('unit', ['karma']);
  grunt.registerTask('staging', ['protractor:staging']);
  grunt.registerTask('localhost', ['protractor:localhost']);
  grunt.registerTask('files', ['concat:appjs']);

  grunt.registerTask('spec', ['files','concat:spec', 'shell:gotest', 'go:build:myapp']);  
  grunt.registerTask('dev', ['files', 'shell:gotest', 'go:build:myapp', 'shell:run']);  
  grunt.registerTask('production', ['files', 'shell:gotest', 'go:build:myapp']);  

  // database tasks
  grunt.registerTask('pgusers', ['shell:psqlusers']);
  grunt.registerTask('pgcreate', ['shell:psqlcreate']);
  grunt.registerTask('pgdrop', ['shell:psqldrop']);
  grunt.registerTask('pgdump', ['shell:psqldump']);
  grunt.registerTask('pglist', ['shell:psqllist']);
  grunt.registerTask('pgrebuild', ['shell:psqldrop', 'shell:psqlcreate']);
  grunt.registerTask('pgrestore', ['shell:psqlrestore']);

  // JS distribution task.   , 'uglify:core'    , 'commonjs'
  grunt.registerTask('dist-js', ['concat']);

  // CSS distribution task.
  grunt.registerTask('less-compile', ['less:compileCore', 'less:compileTheme']);
  grunt.registerTask('dist-css', ['less-compile', 'autoprefixer:core', 'autoprefixer:theme', 'cssmin:minifyTheme']);

  // Full distribution task.
  grunt.registerTask('files', ['clean', 'dist-css', 'dist-js', 'copy:dist_webapp_css', 'copy:dist_webapp_js', 'copy:fonts', 'concat:appjs']);
  //'copy:dist_webapps', 'copy:fonts',

  // default build is dev
  grunt.registerTask('build', ['dev']);
  grunt.registerTask('test', ['e2e', 'unit']);
  grunt.registerTask('run', ['build', 'shell:run']);
  grunt.registerTask('got', ['shell:gotest']);
  grunt.registerTask('pack', ['compress:main']);

};
