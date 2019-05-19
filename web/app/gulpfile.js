var gulp = require('gulp');
var sass = require('gulp-sass');
var cleanCSS = require('gulp-clean-css');
var minify = require('gulp-minify');
var concat = require('gulp-concat');
var clean = require('gulp-clean');
var htmlmin = require('gulp-htmlmin');
var fileinclude = require('gulp-file-include');

gulp.task('copyfavicon', function () {
    return gulp.src('assets/favicon.png')
        .pipe(gulp.dest('../static/'));
});

gulp.task('copyFontelloFonts', function () {
    return gulp.src('vendor/fontello/font/*')
        .pipe(gulp.dest('../static/font'));
});

gulp.task('copyFiles', gulp.series(["copyfavicon","copyFontelloFonts"]));

gulp.task('sass', function(){
    return gulp.src('scss/**/*.scss')
        .pipe(sass())
        .pipe(cleanCSS({compatibility: 'ie8'}))
        .pipe(gulp.dest('../static/css'))
});

gulp.task('js-fe26', function() {
    return gulp.src('js/fe26/fe26.js')
        .pipe(fileinclude({
            prefix: '@@',
            basepath: '@file'
        }))
        .pipe(minify())
        .pipe(gulp.dest('../static/js'))
});

gulp.task('html', function () {
    return gulp.src('html/*.html')
        .pipe(fileinclude({
            prefix: '@@',
            basepath: '@file'
        }))
        .pipe(htmlmin({ collapseWhitespace: true }))
        .pipe(gulp.dest('../templates'));
});

gulp.task('clean', function () {
    return gulp.src(['../static/css','../static/js','../templates','../static/font',"../static/favicon.png"], {read: false,allowEmpty: true})
        .pipe(clean({force: true}));
});

gulp.task('build', gulp.series(["clean","sass",'js-fe26',"html","copyFiles"]));

gulp.task("watch",gulp.series(["clean","sass",'js-fe26',"html","copyFiles",function () {
    gulp.watch('scss/**/*.scss', gulp.series("sass"));
    gulp.watch(['js/fe26/**/*.js'], gulp.series('js-fe26'));
    gulp.watch(['html/**/*.html'], gulp.series("html"));
}]));





