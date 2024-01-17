<?php

use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider and all of them will
| be assigned to the "web" middleware group. Make something great!
|
*/
use App\Http\Controllers\CustomerStoreController;
Route::get('/', function () {
    return view('welcome');
    
});
Route::get('/customer-store', [CustomerStoreController::class, 'getListCustomer']);
Route::get('/customer-store/{id}', [CustomerStoreController::class, 'getDetailCustomer']);
Route::post('/customer-store', [CustomerStoreController::class, 'addCustomerStore'])->withoutMiddleware(['web', 'auth']);
Route::delete('/customer-store/{id}', [CustomerStoreController::class, 'deleteCustomerStore'])->withoutMiddleware(['web', 'auth']);
Route::patch('/customer-store/{id}', [CustomerStoreController::class, 'updateCustomerStore'])->withoutMiddleware(['web', 'auth']);
Route::post('/address/{id}', [CustomerStoreController::class, 'addAddress'])->withoutMiddleware(['web', 'auth']);
Route::patch('/address/{id}', [CustomerStoreController::class, 'updateAddress'])->withoutMiddleware(['web', 'auth']);
Route::delete('/address/{id}', [CustomerStoreController::class, 'deleteAddress'])->withoutMiddleware(['web', 'auth']);