'use strict';
module.exports = function (grunt) {

    grunt.loadNpmTasks('grunt-contrib-clean');
    grunt.loadNpmTasks('grunt-contrib-concat');
    grunt.loadNpmTasks('grunt-contrib-connect');
    grunt.loadNpmTasks('grunt-contrib-copy');
    grunt.loadNpmTasks('grunt-contrib-jshint');
    grunt.loadNpmTasks('grunt-contrib-less');
    grunt.loadNpmTasks('grunt-contrib-uglify');
    grunt.loadNpmTasks('grunt-contrib-watch');
    grunt.loadNpmTasks('grunt-html2js');
    grunt.loadNpmTasks('grunt-karma');

    require('load-grunt-tasks')(grunt);

    // Time how long tasks take. Can help when optimizing build times
    require('time-grunt')(grunt);

    // Default task.
    grunt.registerTask('build', ['clean:build', 'html2js', 'concat',  /* 'less:dist',*/ 'copy:assets', 'copy:build', 'connect:livereload', 'watch']);
    grunt.registerTask('build_without_watch', ['clean', 'html2js', 'concat',/* 'less:dist',*/ 'copy:assets', 'copy:build']);
    grunt.registerTask('release', ['clean:release', 'html2js',/* 'jshint' ,*/ 'concat',/*'karma:unit', 'less:dist',*/ 'copy:assets', 'copy:release']);
    grunt.registerTask('test-watch', ['karma:watch']);

    // Print a timestamp (useful for when watching)
    grunt.registerTask('timestamp', function () {
        grunt.log.subhead(Date());
    });

    var karmaConfig = function (configFile, customOptions) {
      var options = { configFile: configFile, keepalive: true };
      return grunt.util._.extend(options, customOptions);
    };

    // Project configuration.
    grunt.initConfig({
        distdir: 'local_serv',
        releaseDir: '../../../www',
        pkg: grunt.file.readJSON('package.json'),
        project: {
            app: require('../../bower.json').appPath || 'client',
        },
        banner: '/*! <%= pkg.title || pkg.name %> - v<%= pkg.version %> - <%= grunt.template.today("yyyy-mm-dd") %>\n' +
            '<%= pkg.homepage ? " * " + pkg.homepage + "\\n" : "" %>' +
            ' * Copyright (c) <%= grunt.template.today("yyyy") %> <%= pkg.author %>;\n' +
            ' * Licensed <%= _.pluck(pkg.licenses, "type").join(", ") %>\n */\n',
        src: {
            js: [
                'src/**/**/*.js',
                '../common/**/**/*.js'
            ],
            jsTpl: ['<%= distdir %>/templates/**/*.js'],
            specs: ['test/**/*.spec.js'],
            html: [
                'src/index.html'
            ],
            tpl: {
                app: ['src/app/**/*.tpl.html'],
                common: ['../common/**/*.tpl.html'],
                bootstrap: ['../common/bootstrap/template/**/*.html']
            },
            less: ['src/less/stylesheet.less']
        },
        clean: {
          build: {
            options: {force: true},
            src: ['<%= distdir %>/*']
          },
          release: {
            options: {force: true},
            src: ['<%= releaseDir %>/*.js']
          }
        },
        // TODO rewrite
        html2js: {
            app: {
                options: {
                    base: 'src/app'
                },
                src: ['<%= src.tpl.app %>'],
                dest: '<%= distdir %>/templates/app.js',
                module: 'templates.user'
            },
            common: {
                options: {
                    base: '../../source/common'
                },
                src: ['<%= src.tpl.common %>'],
                dest: '<%= distdir %>/templates/common.js',
                module: 'templates.common'
            },
            bootstrap: {
                options: {
                    base: '../../source/common/bootstrap'
                },
                src: ['<%= src.tpl.bootstrap %>'],
                dest: '<%= distdir %>/templates/bootstrap.js',
                module: 'templates.bootstrap'
            }
        },
        concat: {
            dist: {
                options: {
                    banner: "<%= banner %>"
                },
                src: ['<%= src.js %>', '<%= src.jsTpl %>'],
                dest: '<%= distdir %>/<%= pkg.name %>.js'
            },
            angular: {
                src: [
                    '../../vendor/ng-file-upload/angular-file-upload-html5-shim.min.js',
                    '../../vendor/angular/angular.min.js',
//                    'vendor/angular/angular.js',

                    '../../vendor/angular-ui-router/release/angular-ui-router.min.js',
                    '../../vendor/angular-i18n/angular-locale_ru-ru.js',
                    '../../vendor/angular-sanitize/angular-sanitize.min.js',
                    '../../vendor/angular-animate/angular-animate.min.js',
                    '../../vendor/angular-resource/angular-resource.min.js',
                    '../../vendor/angular-bootstrap/ui-bootstrap.min.js',
                    '../../vendor/restangular/dist/restangular.min.js',
                    '../../vendor/ng-file-upload/angular-file-upload.min.js',
                    '../../vendor/angular-redactor/angular-redactor.js',
                    '../../vendor/angular-ui-sortable/sortable.min.js',
                    '../../vendor/angular-ui-utils/ui-utils.min.js',
                    '../../vendor/angular-xeditable/dist/js/xeditable.min.js',
                    //'../../vendor/angular-mocks/angular-mocks.js',
                    '../../vendor/angular-cookies/angular-cookies.min.js',
                    '../../vendor/angular-ui-slider/src/slider.js',
                    '../../vendor/ng-table/ng-table.js',

                    '../libs/angular-dialog-service/dialogs.js',
                    '../libs/angular-ui-select2.js',

                    // Правильный datetime picker со своими зависимостями :)
                    // Всё подключать нельзя так как $modal сдохнет
                    '../../vendor/angular-strap-sass/dist/modules/datepicker.min.js',
                    '../../vendor/angular-strap-sass/dist/modules/datepicker.tpl.min.js',
                    '../../vendor/angular-strap-sass/dist/modules/timepicker.min.js',
                    '../../vendor/angular-strap-sass/dist/modules/timepicker.tpl.min.js',
                    '../../vendor/angular-strap-sass/dist/modules/date-parser.min.js',
                    '../../vendor/angular-strap-sass/dist/modules/tooltip.min.js',
                    '../../vendor/angular-strap-sass/dist/modules/tooltip.tpl.min.js',
                    '../../vendor/angular-strap-sass/dist/modules/dimensions.min.js'

                ],
                dest: '<%= distdir %>/angular.js'
            },

            jquery: {
                src: [
                    '../../vendor/jquery/dist/jquery.min.js',
                    '../../vendor/jquery-ui/ui/minified/jquery.ui.core.min.js',
                    '../../vendor/jquery-ui/ui/minified/jquery.ui.widget.min.js',
                    '../../vendor/jquery-ui/ui/minified/jquery.ui.mouse.min.js',
                    '../../vendor/jquery-ui/ui/minified/jquery.ui.sortable.min.js',
                    '../../vendor/jquery-ui/ui/minified/jquery.ui.slider.min.js',
                    '../libs/jquery-ui-touch-punch/jquery.ui.touch-punch.min.js',
                    '../../vendor/lodash/dist/lodash.underscore.min.js',
                    '../libs/redactor/js/redactor.min.js',
                    '../libs/redactor/locales/redactor.ru.js',
                    '../../vendor/select2/select2.min.js'
//                  'vendor/pace/pace.min.js'
                ],
                dest: '<%= distdir %>/jquery.js'
            }
        },
        copy: {
            assets: {
                files: [
                    { dest: '<%= distdir %>', src: ['**/*.png', '**/*.gif'], expand: true, cwd: 'assets/' },     // common assets
                    { dest: '<%= distdir %>', src: ['**/*.png', '**/*.gif'], expand: true, cwd: '../../source/assets/' },
                    { dest: '<%= distdir %>/fonts/', src: ['*.woff', '*.ttf'], expand: true, cwd: '../../vendor/bootstrap/fonts/' }, // bootstrap
                    { expand: true, cwd: '../../source/assets/', dest: '<%= distdir %>/', src: 'favicon.ico'},
                ]
            },
            build: {
                files: [
                    { expand: true, cwd: 'src', dest: '<%= distdir %>/', src: 'index.html'}
                ]
            },
            release: {
               files: [
                  { expand: true, cwd: '<%= distdir %>/', dest: '<%= releaseDir %>/', src: 'jquery.js'},
                  { expand: true, cwd: '<%= distdir %>/', dest: '<%= releaseDir %>/', src: 'angular.js'},
                  { expand: true, cwd: '<%= distdir %>/', dest: '<%= releaseDir %>/', src: 'zegota.js'},
                  { expand: true, cwd: '<%= distdir %>/', dest: '<%= releaseDir %>/', src: 'zegota.css'},
                  { expand: true, cwd: '<%= distdir %>/', dest: '<%= releaseDir %>/', src: 'zegota.css.map'}
               ]
            }
        },
        less: {
            dist: {
                options: {
                    paths: ["src/less"],
                    compress: true,
                    cleancss: true,
                    optimization: 2,
                    sourceMap: true,
                    sourceMapFilename: '<%= distdir %>/<%= pkg.name %>.css.map',
                    sourceMapBasepath: '<%= distdir %>/'
                },
                files: {
                    '<%= distdir %>/<%= pkg.name %>.css': 'src/less/stylesheet.less'
                }
            }
        },
        watch: {
            all: {
                files: ['<%= src.js %>', '<%= src.less %>', '<%= src.tpl.app %>', '<%= src.tpl.common %>', '<%= src.tpl.bootstrap %>', '<%= src.html %>'],
                tasks: ['build_without_watch', 'timestamp'],
                options: {
                    livereload: true
                }
            },
            livereload: {
                options: {
                    livereload: '<%= connect.options.livereload %>'
                },
                files: [
                    'gruntFile.js',
                    '<%= project.app %>/{,*/}*.html',
                    '<%= project.app %>/images/{,*/}*.{png,jpg,jpeg,gif,webp,svg}',
                    '<%= project.app %>/img/{,*/}*.{png,jpg,jpeg,gif,webp,svg}',
                    '../templates/{,*/}*.{js}'
                ]
            }
        },
        connect: {
            options: {
                port: 9000,
                hostname: '0.0.0.0',
                livereload: 35729
            },
            livereload: {
                options: {
                    open: true,
                    base: [
                        '<%= distdir %>'
                    ]
                }
            }
        },
        karma: {
          unit: {
            options: karmaConfig('test/config/unit.conf.js')
          },
          watch: {
            options: karmaConfig('test/config/unit.conf.js', {singleRun: false, autoWatch: true})
          }
        },
        jshint: {
          files: ['gruntFile.js', '<%= src.js %>', '<%= src.jsTpl%>', '<%= src.specs %>'],
          options:{
            curly: true,
            eqeqeq: true,
            immed: true,
            latedef: true,
            newcap: true,
            noarg: true,
            sub: true,
            boss: true,
            eqnull: true,
            globals: {}
          }
        }
    });
};
