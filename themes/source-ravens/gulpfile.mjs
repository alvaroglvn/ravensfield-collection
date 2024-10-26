import gulp from 'gulp';
const { series, parallel, src, dest, watch } = gulp;
import browserSync from 'browser-sync';
const bs = browserSync.create();

import postcss from 'gulp-postcss';
import zip from 'gulp-zip';
import concat from 'gulp-concat';
import terser from 'gulp-terser';
import autoprefixer from 'autoprefixer';
import cssnano from 'cssnano';
import postcssImport from 'postcss-import';
import fs from 'fs';

// CSS task
function css() {
    return src('assets/css/screen.css', { sourcemaps: true })
        .pipe(postcss([postcssImport(), autoprefixer(), cssnano()]))
        .pipe(dest('assets/built/', { sourcemaps: '.' }))
        .pipe(bs.stream());
}

// JavaScript task
function js() {
    return src(['assets/js/lib/*.js', 'assets/js/*.js'], { sourcemaps: true })
        .pipe(concat('source.js'))
        .pipe(terser())
        .pipe(dest('assets/built/', { sourcemaps: '.' }))
        .pipe(bs.stream());
}

// Handlebars (HBS) templates task
function hbs() {
    return src(['*.hbs', 'partials/**/*.hbs']).pipe(bs.stream());
}

// Zip task for packaging the theme
function zipper() {
    const filename = JSON.parse(fs.readFileSync('./package.json')).name + '.zip';
    return src(
        ['**', '!node_modules/**', '!dist/**', '!yarn-error.log', '!yarn.lock', '!gulpfile.mjs'],
        { nodir: true }
    )
        .pipe(zip(filename))
        .pipe(dest('dist/'));
}

// Serve task with BrowserSync for live-reloading
function serve(done) {
    bs.init({
        server: { baseDir: './dist' },
        port: 3000,
        open: false,
        notify: false,
    });
    done();
}

// Watcher tasks for development
const cssWatcher = () => watch('assets/css/**', css);
const jsWatcher = () => watch('assets/js/**', js);
const hbsWatcher = () => watch(['*.hbs', 'partials/**/*.hbs'], hbs);

// Main build task to compile CSS and JS
const build = series(css, js);

// Default task for development: build and watch files
const watcher = parallel(cssWatcher, jsWatcher, hbsWatcher);
export default series(build, serve, watcher);

// Export build and zip tasks
export { build, zipper as zip };