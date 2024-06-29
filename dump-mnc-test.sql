PGDMP  ,                    |            mnc-test    16.2    16.2     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    16556    mnc-test    DATABASE     �   CREATE DATABASE "mnc-test" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_Indonesia.1252';
    DROP DATABASE "mnc-test";
                postgres    false                        2615    2200    public    SCHEMA        CREATE SCHEMA public;
    DROP SCHEMA public;
                pg_database_owner    false            �           0    0    SCHEMA public    COMMENT     6   COMMENT ON SCHEMA public IS 'standard public schema';
                   pg_database_owner    false    5            �            1259    16659    balances    TABLE       CREATE TABLE public.balances (
    balance_id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid,
    balance_before numeric,
    balance_after numeric,
    created_date timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    status character varying(1)
);
    DROP TABLE public.balances;
       public         heap    postgres    false    5    5    5            �           0    0    COLUMN balances.status    COMMENT     F   COMMENT ON COLUMN public.balances.status IS 'L = Latest, O = Oldest';
          public          postgres    false    217            �            1259    16668    transactions    TABLE     �  CREATE TABLE public.transactions (
    transaction_id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid,
    balance_id uuid,
    transaction_type character varying(1),
    amount numeric,
    status character varying(1),
    transaction_category character varying(2),
    created_date timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    remarks character varying,
    target_user character varying
);
     DROP TABLE public.transactions;
       public         heap    postgres    false    5    5    5            �           0    0 $   COLUMN transactions.transaction_type    COMMENT     R   COMMENT ON COLUMN public.transactions.transaction_type IS 'D =  DebitC = Credit';
          public          postgres    false    218            �           0    0    COLUMN transactions.status    COMMENT     X   COMMENT ON COLUMN public.transactions.status IS 'S = Success, F = Failed, P = Pending';
          public          postgres    false    218            �           0    0 (   COLUMN transactions.transaction_category    COMMENT     j   COMMENT ON COLUMN public.transactions.transaction_category IS 'TP = Top Up, PY = Payment, TF = Transfer';
          public          postgres    false    218            �            1259    16648    users    TABLE     *  CREATE TABLE public.users (
    user_id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    first_name text,
    last_name text,
    phone_number text,
    pin text,
    created_date timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    address text,
    update_date timestamp with time zone
);
    DROP TABLE public.users;
       public         heap    postgres    false    5    5    5            �          0    16659    balances 
   TABLE DATA           l   COPY public.balances (balance_id, user_id, balance_before, balance_after, created_date, status) FROM stdin;
    public          postgres    false    217   �       �          0    16668    transactions 
   TABLE DATA           �   COPY public.transactions (transaction_id, user_id, balance_id, transaction_type, amount, status, transaction_category, created_date, remarks, target_user) FROM stdin;
    public          postgres    false    218   �       �          0    16648    users 
   TABLE DATA           v   COPY public.users (user_id, first_name, last_name, phone_number, pin, created_date, address, update_date) FROM stdin;
    public          postgres    false    216   �       7           2606    16667    balances balances_pkey 
   CONSTRAINT     \   ALTER TABLE ONLY public.balances
    ADD CONSTRAINT balances_pkey PRIMARY KEY (balance_id);
 @   ALTER TABLE ONLY public.balances DROP CONSTRAINT balances_pkey;
       public            postgres    false    217            9           2606    16675    transactions transactions_pkey 
   CONSTRAINT     h   ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (transaction_id);
 H   ALTER TABLE ONLY public.transactions DROP CONSTRAINT transactions_pkey;
       public            postgres    false    218            3           2606    16658    users uni_users_phone_number 
   CONSTRAINT     _   ALTER TABLE ONLY public.users
    ADD CONSTRAINT uni_users_phone_number UNIQUE (phone_number);
 F   ALTER TABLE ONLY public.users DROP CONSTRAINT uni_users_phone_number;
       public            postgres    false    216            5           2606    16656    users users_pkey 
   CONSTRAINT     S   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            postgres    false    216            �   H  x���1r$1E��Sl��0gp�/�D��G�fg�	:R$�*v���W�����-Ǣ�C�R{�^�l2T
gw�5�8�,B�������*`�h���˴8���s���g��͕���85kk��#�NBж�^��IsX���Cn�:�I���P[g0��u�*I�v���h����ȓ�}�M��X�|K�,�c�܅�.z���Y�VKN7X��&`1��u�5����)a}�]N����u˖
֗@��k؃�����)�y�#�I�H��^\��!=c��wvmB��@�zg+I�n��������zqۭˋ�{������      �   �  x�����1��Oq{�c�Ǿ�"ꕠA���Db�(x{2�m��j�p1J�q��s�Zł ��5H�Vg�֔���K6+��d0dreU�9$)�H����s��lN�C
�����Luv�6£��ۇǵ�X��rK�F�+!�7ض����R��N	����	��ܡ�l�Z`=�/��e�R2��g����I���Ϳ�?��ˍ�Ε����Ӟ�������\G'�ޢB��J���L�6MqNH�&Q���,�#�ֆQ�%��׍8�ة(�wb�L�XC�@�
tw���v2I��\��i9�N�Pc�>���r������.-)�y~YX԰C���Y�q(h��Yj*��#��y~�a߯O�d_�߾�95��;83tt\��	�g�늎�ϩ��HR��T�Ap,�	%SYq���񞄰�USsK�� ��5�5P[oK�4���}2[�Լ��{5��?���D;s      �   �   x���;�0��ڜ�-�]��gIc��"R��"$�4���t��W�$6K*��dS"�[�d{e'潷�Ƕoϻ���O�F��>;�)�hyB0G���l��*:]g�U� ,|o�`�҄)���[a�*�*�j������m��3�7T     