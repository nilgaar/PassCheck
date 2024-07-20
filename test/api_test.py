import requests
 
def test_check_hash_nohit():
    no_results = requests.get('http://localhost:8080?hash=0654804158fd3078157c3dce72f60424e275bd300ce1d1f8433124ec8dd8cbe1125ce8a1d5efd585b0546707adf171a8f98c014ba8b4f5a09518c6b063d3f057')
    assert no_results.status_code == 200
    assert no_results.json() == {'data': 'No matches found'}
    
    
def test_check_hash_hit():
    results = requests.get('http://localhost:8080?hash=0430c0ebb2398fda098d1acc0372de9b4c73a0e52bad59385a691cf61520f7dcb808fd75789d7094dbf4b452c3b1c9bc8f741cd2cbde2ec354b34a22e22be745')
    assert results.status_code == 200
    
    file_list = results.json()['data'].split(", ")

    # file_list must contain: "500-worst-passwords.txt, best1050.txt, 10-million-password-list-top-1000.txt, 10k-most-common.txt, 10-million-password-list-top-500.txt, 10-million-password-list-top-10000.txt, 100k-most-used-passwords-NCSC.txt, 10-million-password-list-top-100000.txt, 10-million-password-list-top-1000000.txt"

    expected_files = [
        "10-million-password-list-top-500.txt",
        "10-million-password-list-top-1000.txt",
        "10-million-password-list-top-10000.txt",
        "10-million-password-list-top-100000.txt",
        "10-million-password-list-top-1000000.txt",
        "500-worst-passwords.txt",
        "best1050.txt",
        "10k-most-common.txt",
        "100k-most-used-passwords-NCSC.txt"
    ]

    for file in expected_files:
        assert file in file_list

    
def test_single_hit():
    results = requests.get('http://localhost:8080?hash=002e136b5760c542271a09cf623a62fd990a23694cf99a5c9306e09c46d8ad6b14100606dd56bf2053626e1a94ccbca54a0b534d695abc0348a396a379677b61')
    assert results.status_code == 200
    assert results.json() == {'data': '1900-2020.txt'}