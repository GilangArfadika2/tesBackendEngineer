<?php

namespace App\Http\Controllers;

use App\Models\CustomerStore;
use App\Models\Address;
use Illuminate\Http\Request;
use Illuminate\Validation\Rule;

class CustomerStoreController extends Controller
{
    public function getListCustomer()
    {
        $customerStores = CustomerStore::orderBy('title')->get();

        return response()->json($customerStores, 200);
    }

    public function getDetailCustomer($id)
    {
        $customerAndAddress = CustomerStore::with(['address' => function ($query) {
            $query->orderBy('address', 'asc');}])->find($id);

        if ($customerAndAddress) {
            return response()->json($customerAndAddress, 200);
        } else {
            return response()->json([
                'message' => 'Customer Store tidak ditemukan',
                'status' => 404,
            ], 404);
        }
    }

    // Method Post Customer Store
    public function addCustomerStore(Request $request)
    {
        // periksa validasi masukan
        try {
            $validatedData = $request->validate([
                'title' => 'required|string|max:255',
                'name' => 'required|string|max:255',
                'gender' => 'required|string|in:M,F',
                'phone_number' => 'required|string|max:20',
                'image' => 'nullable|string|max:255',
                'email' => 'required|email|unique:customer_stores,email',
            ]);
        
            $customerStore = CustomerStore::simpan($validatedData);

            return response()->json($customerStore, 201);
        }

        catch (\Illuminate\Validation\ValidationException $e) {
            //jika validasi gagal 
            return response()->json([
                'message' => 'Validasi gagal',
                'errors' => $e->errors(),
            ], 422);
        }
    }
    // Method Post Address
    public function addAddress(Request $request, $id)
    {

               
        $customerStore = CustomerStore::find($id);
        // Periksa apakah Customer Store Ada
        if ($customerStore){
            // tes validasi masukan
            try {
                $validatedData = $request->validate([
                    'address' => 'required|string|max:255',
                    'district' => 'required|string|max:255',
                    'city' => 'required|string|max:255',
                    'province' => 'required|string|max:255',
                    'postal_code' => 'required|integer',
                ]);
                $address =  Address:: simpan($validatedData, $id);

                return response()->json($address, 201);
            }

            catch (\Illuminate\Validation\ValidationException $e) {
                //jika validasi gagal 
                return response()->json([
                    'message' => 'Validasi gagal',
                    'errors' => $e->errors(),
                ], 422);
            }
           
        }
        else {
            return response()->json([
                'message' => 'Customer Store tidak ditemukan',
                'status' => 404,
            ], 404);;
        }
        
     
    }
    // Method delete customer store
    public function deleteCustomerStore($id)
    {
        $customerStore = CustomerStore::with('address')->find($id);

        if ( $customerStore) {
           
            $customerStore->address->each(function ($address) {
                $address->delete();
            });
            $customerStore->delete();
            return response()->json([
                'message' => 'Customer Store Telah dihapus',
                'status' => 200,
            ], 200);
        } else {
            return response()->json([
                'message' => 'Customer Store tidak ditemukan',
                'status' => 404,
            ], 404);;
        }
    }
    // Method hapus Address
    public function deleteAddress($id)
    {
        $address = Address::find($id);

        if ($address) {
            $address->delete();
            return response()->json([
                'message' => 'Address Telah dihapus',
                'status' => 200,
            ], 200);;
        } else {
            return response()->json([
                'message' => 'Address Tidak ditemukan',
                'status' => 404,
            ], 404);;
        }
    }

    
    // Method update customer store
    public function updateCustomerStore(Request $request,$id) {

        try {
            $validatedData = $request->validate([
                'title' => 'string|max:255',
                'name' => 'string|max:255',
                'gender' => 'string|in:M,F',
                'phone_number' => 'string|max:20',
                'image' => 'nullable|string|max:255',
                'email' => [
                    'email',
                    Rule::unique('customer_stores', 'email')->ignore($id),
                ],
            ]);
         }   

         catch (\Illuminate\Validation\ValidationException $e) {
            //jika validasi gagal 
            return response()->json([
                'message' => 'Validasi gagal',
                'errors' => $e->errors(),
            ], 422);
        }

        $customerStore = CustomerStore::find($id);

        if ($customerStore){

            $customerStore->updated_at = now();

            $customerStore->update($validatedData);

            return response()->json($customerStore, 200);

            
        }
        else {
            return response()->json([
                'message' => 'Customer Store tidak ditemukan',
                'status' => 404,
            ], 404);;
        }

    }

    // Method update Address
    public function updateAddress(Request $request,$id) {

        try {
            $validatedData = $request->validate([
                'customer_store_id' => 'exists:customer_stores,id',
                'address' => 'string|max:255',
                'district' => 'string|max:255',
                'city' => 'string|max:255',
                'province' => 'string|max:255',
                'postal_code' => 'integer',
            ]);
        }  
        catch (\Illuminate\Validation\ValidationException $e) {
            //jika validasi gagal 
            return response()->json([
                'message' => 'Validasi gagal',
                'errors' => $e->errors(),
            ], 422);
        }

        $address= Address::find($id);

        if ($address){

            $address->updated_at = now();

            $address->update($validatedData);

            return response()->json($address, 200);

            
        }
        else {
            return response()->json([
                'message' => 'Address tidak ditemukan',
                'status' => 404,
            ], 404);;
        }

    }
}
